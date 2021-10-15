/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */

export interface paths {
  "/userinfo": {
    /** get user info */
    get: operations["getUserInfo"];
  };
  "/groups/{id}/gcp_projects": {
    /** Query and return all GCP projects for the group */
    get: operations["getGCPProjects"];
    parameters: {
      path: {
        /** Group ID */
        id: string;
      };
    };
  };
  "/gcp/{id}/tables": {
    /** Return all BigQuery tables in gcp project */
    get: operations["getBigqueryTables"];
    parameters: {
      path: {
        /** GCP project ID */
        id: string;
      };
    };
  };
  "/collections": {
    /** List all DataproductCollections */
    get: operations["getDataproductCollections"];
    /** Create a new DataproductCollection */
    post: operations["createDataproductCollection"];
  };
  "/collections/{id}": {
    /** List a DataproductCollection with dataproducts */
    get: operations["getDataproductCollection"];
    /** Update a DataproductCollection */
    put: operations["updateDataproductCollection"];
    /** Delete a DataproductCollection */
    delete: operations["deleteDataproductCollection"];
  };
  "/dataproducts": {
    /** Get dataproducts */
    get: operations["getDataproducts"];
    /** Create a new dataproduct */
    post: operations["createDataproduct"];
  };
  "/dataproducts/{id}": {
    /** Get dataproduct */
    get: operations["getDataproduct"];
    /** Update a dataproduct */
    put: operations["updateDataproduct"];
    /** Delete a dataproduct */
    delete: operations["deleteDataproduct"];
  };
  "/dataproducts/{id}/metadata": {
    /** Get dataproduct metadata */
    get: operations["getDataproductMetadata"];
  };
  "/search": {
    /** Search in NADA */
    get: operations["search"];
    parameters: {
      query: {
        q?: string;
        limit?: number;
        offset?: number;
      };
    };
  };
}

export interface components {
  schemas: {
    DataproductCollection: {
      id: string;
      name: string;
      description?: string;
      slug: string;
      repo?: string;
      last_modified: string;
      created: string;
      owner: components["schemas"]["Owner"];
      keywords?: string[];
      dataproducts: components["schemas"]["DataproductSummary"][];
    };
    NewDataproductCollection: {
      name: string;
      description?: string;
      slug?: string;
      repo?: string;
      owner: components["schemas"]["Owner"];
      keywords?: string[];
    };
    UpdateDataproductCollection: {
      name: string;
      description?: string;
      slug?: string;
      repo?: string;
      keywords?: string[];
    };
    Dataproduct: {
      id: string;
      name: string;
      description?: string;
      slug?: string;
      repo?: string;
      pii: boolean;
      owner: components["schemas"]["Owner"];
      type: components["schemas"]["DataproductType"];
      datasource: components["schemas"]["Datasource"];
    };
    Datasource: components["schemas"]["Bigquery"];
    NewDataproduct: {
      name: string;
      description?: string;
      slug?: string;
      repo?: string;
      pii: boolean;
      owner: components["schemas"]["Owner"];
      datasource: components["schemas"]["Datasource"];
    };
    UpdateDataproduct: {
      name: string;
      description?: string;
      slug?: string;
      repo?: string;
      pii: boolean;
    };
    DataproductSummary: {
      id: string;
      name: string;
      type: components["schemas"]["DataproductType"];
    };
    DataproductType: "bigquery";
    Owner: {
      group: string;
      teamkatalogen?: string;
    };
    Bigquery: {
      project_id: string;
      dataset: string;
      table: string;
    };
    UserInfo: {
      name: string;
      email: string;
      groups: string[];
    };
    TableColumn: {
      name: string;
      type: string;
      mode: string;
      description: string;
    };
    DataproductMetadata: {
      dataproduct_id: string;
      schema: components["schemas"]["TableColumn"][];
    };
    SearchResultEntry: {
      url: string;
      type: components["schemas"]["SearchResultType"];
      id: string;
      name: string;
      excerpt: string;
    };
    SearchResultType: "dataproduct" | "DataproductCollection" | "datapackage";
  };
}

export interface operations {
  /** get user info */
  getUserInfo: {
    responses: {
      /** OK */
      200: {
        content: {
          "application/json": components["schemas"]["UserInfo"];
        };
      };
    };
  };
  /** Query and return all GCP projects for the group */
  getGCPProjects: {
    parameters: {
      path: {
        /** Group ID */
        id: string;
      };
    };
    responses: {
      /** OK */
      200: {
        content: {
          "application/json": string[];
        };
      };
    };
  };
  /** Return all BigQuery tables in gcp project */
  getBigqueryTables: {
    parameters: {
      path: {
        /** GCP project ID */
        id: string;
      };
    };
    responses: {
      /** OK */
      200: {
        content: {
          "application/json": components["schemas"]["Bigquery"][];
        };
      };
    };
  };
  /** List all DataproductCollections */
  getDataproductCollections: {
    parameters: {
      query: {
        limit?: number;
        offset?: number;
      };
    };
    responses: {
      /** OK */
      200: {
        content: {
          "application/json": components["schemas"]["DataproductCollection"][];
        };
      };
    };
  };
  /** Create a new DataproductCollection */
  createDataproductCollection: {
    responses: {
      /** Created successfully */
      201: {
        content: {
          "application/json": components["schemas"]["DataproductCollection"];
        };
      };
    };
    requestBody: {
      content: {
        "application/json": components["schemas"]["NewDataproductCollection"];
      };
    };
  };
  /** List a DataproductCollection with dataproducts */
  getDataproductCollection: {
    parameters: {
      path: {
        /** DataproductCollection ID */
        id: string;
      };
    };
    responses: {
      /** OK */
      200: {
        content: {
          "application/json": components["schemas"]["DataproductCollection"][];
        };
      };
    };
  };
  /** Update a DataproductCollection */
  updateDataproductCollection: {
    parameters: {
      path: {
        /** DataproductCollection ID */
        id: string;
      };
    };
    responses: {
      /** Updated OK */
      200: {
        content: {
          "application/json": components["schemas"]["DataproductCollection"];
        };
      };
    };
    requestBody: {
      content: {
        "application/json": components["schemas"]["UpdateDataproductCollection"];
      };
    };
  };
  /** Delete a DataproductCollection */
  deleteDataproductCollection: {
    parameters: {
      path: {
        /** DataproductCollection ID */
        id: string;
      };
    };
    responses: {
      /** Deleted OK */
      204: never;
    };
  };
  /** Get dataproducts */
  getDataproducts: {
    parameters: {
      query: {
        limit?: number;
        offset?: number;
      };
    };
    responses: {
      /** OK */
      200: {
        content: {
          "application/json": components["schemas"]["Dataproduct"][];
        };
      };
    };
  };
  /** Create a new dataproduct */
  createDataproduct: {
    responses: {
      /** Created successfully */
      201: {
        content: {
          "application/json": components["schemas"]["Dataproduct"];
        };
      };
    };
    requestBody: {
      content: {
        "application/json": components["schemas"]["NewDataproduct"];
      };
    };
  };
  /** Get dataproduct */
  getDataproduct: {
    parameters: {
      path: {
        /** Dataproduct ID */
        id: string;
      };
    };
    responses: {
      /** OK */
      200: {
        content: {
          "application/json": components["schemas"]["Dataproduct"];
        };
      };
    };
  };
  /** Update a dataproduct */
  updateDataproduct: {
    parameters: {
      path: {
        /** Dataproduct ID */
        id: string;
      };
    };
    responses: {
      /** Updated OK */
      200: {
        content: {
          "application/json": components["schemas"]["Dataproduct"];
        };
      };
    };
    requestBody: {
      content: {
        "application/json": components["schemas"]["UpdateDataproduct"];
      };
    };
  };
  /** Delete a dataproduct */
  deleteDataproduct: {
    parameters: {
      path: {
        /** Dataproduct ID */
        id: string;
      };
    };
    responses: {
      /** Deleted OK */
      204: never;
    };
  };
  /** Get dataproduct metadata */
  getDataproductMetadata: {
    parameters: {
      path: {
        /** Dataproduct ID */
        id: string;
      };
    };
    responses: {
      /** OK */
      200: {
        content: {
          "application/json": components["schemas"]["DataproductMetadata"];
        };
      };
    };
  };
  /** Search in NADA */
  search: {
    parameters: {
      query: {
        q?: string;
        limit?: number;
        offset?: number;
      };
    };
    responses: {
      /** Search result */
      200: {
        content: {
          "application/json": components["schemas"]["SearchResultEntry"][];
        };
      };
    };
  };
}

export interface external {}
