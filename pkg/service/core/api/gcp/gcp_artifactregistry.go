package gcp

import (
	"context"
	"github.com/navikt/nada-backend/pkg/artifactregistry"
	"github.com/navikt/nada-backend/pkg/errs"
	"github.com/navikt/nada-backend/pkg/service"
	"github.com/rs/zerolog"
)

var _ service.ArtifactRegistryAPI = &artifactRegistryAPI{}

type artifactRegistryAPI struct {
	ops artifactregistry.Operations
	log zerolog.Logger
}

func (a *artifactRegistryAPI) AddArtifactRegistryPolicyBinding(ctx context.Context, id *service.ContainerRepositoryIdentifier, binding *service.Binding) error {
	const op errs.Op = "gcp.AddArtifactRegistryPolicyBinding"

	err := a.ops.AddArtifactRegistryPolicyBinding(ctx, &artifactregistry.ContainerRepositoryIdentifier{
		Project:    id.Project,
		Location:   id.Location,
		Repository: id.Repository,
	}, &artifactregistry.Binding{
		Role:    binding.Role,
		Members: binding.Members,
	})
	if err != nil {
		return errs.E(op, err)
	}

	return nil
}

func (a *artifactRegistryAPI) RemoveArtifactRegistryPolicyBinding(ctx context.Context, id *service.ContainerRepositoryIdentifier, binding *service.Binding) error {
	const op errs.Op = "gcp.RemoveArtifactRegistryPolicyBinding"

	err := a.ops.RemoveArtifactRegistryPolicyBinding(ctx, &artifactregistry.ContainerRepositoryIdentifier{
		Project:    id.Project,
		Location:   id.Location,
		Repository: id.Repository,
	}, &artifactregistry.Binding{
		Role:    binding.Role,
		Members: binding.Members,
	})
	if err != nil {
		return errs.E(op, err)
	}

	return nil
}

func (a *artifactRegistryAPI) ListContainerImagesWithTag(ctx context.Context, id *service.ContainerRepositoryIdentifier, tag string) ([]*service.ContainerImage, error) {
	const op errs.Op = "gcp.ListContainerImagesWithTag"

	raw, err := a.ops.ListContainerImagesWithTag(ctx, &artifactregistry.ContainerRepositoryIdentifier{
		Project:    id.Project,
		Location:   id.Location,
		Repository: id.Repository,
	}, tag)
	if err != nil {
		return nil, errs.E(op, err)
	}

	var images []*service.ContainerImage
	for _, image := range raw {
		labels := map[string]string{}

		// Fetch the manifest to get the labels, for now, we ignore errors
		manifest, err := a.ops.GetContainerImageManifest(ctx, image.URI)
		if err != nil {
			a.log.Error().Err(err).Msgf("failed to get manifest for image %s", image.URI)
		}

		if err == nil {
			labels = manifest.Labels
		}

		images = append(images, &service.ContainerImage{
			Name: image.Name,
			URI:  image.URI,
			Manifest: &service.Manifest{
				Labels: labels,
			},
		})
	}

	return images, nil
}

func NewArtifactRegistryAPI(ops artifactregistry.Operations, log zerolog.Logger) service.ArtifactRegistryAPI {
	return &artifactRegistryAPI{
		ops: ops,
		log: log,
	}
}
