import { LogoProps } from './iconBox'

const BigQueryLogo = (logoProps: LogoProps) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={logoProps.size}
    height={logoProps.size}
    viewBox={'0 0 64 64'}
  >
    <path
      d="M14.48 58.196L.558 34.082c-.744-1.288-.744-2.876 0-4.164L14.48 5.805c.743-1.287 2.115-2.08 3.6-2.082h27.857c1.48.007 2.845.8 3.585 2.082l13.92 24.113c.744 1.288.744 2.876 0 4.164L49.52 58.196c-.743 1.287-2.115 2.08-3.6 2.082H18.07c-1.483-.005-2.85-.798-3.593-2.082z"
      fill="#4386fa"
    />
    <path
      d="M40.697 24.235s3.87 9.283-1.406 14.545-14.883 1.894-14.883 1.894L43.95 60.27h1.984c1.486-.002 2.858-.796 3.6-2.082L58.75 42.23z"
      opacity=".1"
    />
    <path
      d="M45.267 43.23L41 38.953a.67.67 0 0 0-.158-.12 11.63 11.63 0 1 0-2.032 2.037.67.67 0 0 0 .113.15l4.277 4.277a.67.67 0 0 0 .947 0l1.12-1.12a.67.67 0 0 0 0-.947zM31.64 40.464a8.75 8.75 0 1 1 8.749-8.749 8.75 8.75 0 0 1-8.749 8.749zm-5.593-9.216v3.616c.557.983 1.363 1.803 2.338 2.375v-6.013zm4.375-2.998v9.772a6.45 6.45 0 0 0 2.338 0V28.25zm6.764 6.606v-2.142H34.85v4.5a6.43 6.43 0 0 0 2.338-2.368z"
      fill="#fff"
    />
  </svg>
)

export default BigQueryLogo
