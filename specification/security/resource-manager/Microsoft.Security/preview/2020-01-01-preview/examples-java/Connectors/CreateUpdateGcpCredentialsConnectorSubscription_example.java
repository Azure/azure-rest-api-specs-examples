import com.azure.resourcemanager.security.models.AutoProvision;
import com.azure.resourcemanager.security.models.GcpCredentialsDetailsProperties;
import com.azure.resourcemanager.security.models.HybridComputeSettingsProperties;

/** Samples for Connectors CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/Connectors/CreateUpdateGcpCredentialsConnectorSubscription_example.json
     */
    /**
     * Sample code: gcpCredentials - Create a cloud account connector for a subscription.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void gcpCredentialsCreateACloudAccountConnectorForASubscription(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .connectors()
            .define("gcp_dev")
            .withHybridComputeSettings(new HybridComputeSettingsProperties().withAutoProvision(AutoProvision.OFF))
            .withAuthenticationDetails(
                new GcpCredentialsDetailsProperties()
                    .withOrganizationId("AscDemoOrg")
                    .withType("service_account")
                    .withProjectId("asc-project-1234")
                    .withPrivateKeyId("6efg587hra2568as34d22326b044cc20dc2af")
                    .withPrivateKey(
                        "-----BEGIN PRIVATE KEY-----\n"
                            + "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCpxYHcLzcDZ6/Q\n"
                            + "AeQZnQXM5GTb3p09Xsbjo2T2F61b6I7FZiQXBrbw3Zf0CUCkkqTTpD5xifl82yQ6\n"
                            + "89V7SAe8hxI7esAcVDhm/aJMqzVjHLISAU2L3li1sn0jjY2oYtndwN6bRivP8O6t\n"
                            + "9F+W6E0zMlbCxtpZEHLbb6WxlJJrwEQ0MPH2yOCwZUQi6NHksAtEzX2nNKJNyUC7\n"
                            + "QyBVHHMm34H2bmZwsuQp3y2otpcJ9tJnVmYfC3k/w4x2L+DIK7JnQP/C1wQqu2du\n"
                            + "c0w6sydF6RhLoHButrVdYRJTdfK4k03SsSTyMqZ+f7LNnKw3xenzw1VmEpk8mvoQ\n"
                            + "t08tCBOrAgMBAAECggEAByzz6iyMtLYjNjV+QJ7kad6VbL2iA8AHxANZ9xTVHPdd\n"
                            + "YXaJu/dqsA+NpqDlfI8+LDva782XH/HbPCqmMUnAGfXTjXQIvqnIoIHD5F2wKfpC\n"
                            + "hIRNlMXXFgbvRxtqi11yO+80+XcjzuwuCmgzyhsTeEB+bkkdXXpWgHPdmv3emnM6\n"
                            + "MQM9Zgrug0UndPmiUwKOcJSU4PlmlTpHEV4vA6JfA4bvphy9m1jxO5qWeah5yym2\n"
                            + "6FP5BRIDF98kFrDnSXJjajwgLCQ+MypFQXyax6XkxDxuKXbng1bv7eZDjqazIChk\n"
                            + "m0y14X0s0jnWc+AX8vfeSf7d+EsGdVinEwR1aAawEQKBgQDqDB0qxcIQ1oI1Kww8\n"
                            + "9vXefTiuWsf47F+fJ/DIOEbiRfE8IdCgmOABvcqJIoxW/DFMBEdLCcx73Km7pOmd\n"
                            + "Kg1ddScnaO8cOj2v/Ub+fAqVrA4ki4ViYP0A7/Nogga3Jr/x3ey5bitrIfFImteS\n"
                            + "CgBHBzZvoQpvO4lB2tKVgo2P9wKBgQC5sgTEq4sasRGSAY6lIoJno0I8w28a/16D\n"
                            + "es60XQeY1ger8uTGwlT02v/u/arDUmRLPClpujXq6gK29KvtRCHy7JkpGbqW2bZs\n"
                            + "PFKKWR7Tk3XPKYyjv94AIi5/xoFeDhS4lpAvy3Z5tQhYS6wqWKvT6yZQ3kM+Hfxs\n"
                            + "pHgvu3mU7QKBgQC9/E1k3hj1cBtMK4CIsHPPQljTd4+iacYJPPPAo6YuoVX8WPqw\n"
                            + "ksgrwbN59Fh1d8xQh5yTtgWOegYx8uFMGcm1lpbM7+pBQKm4hWGuzGQPMRZd5f/F\n"
                            + "ZzOZIi61I+9tlv/yxxIVR+/ozCm/pSneO04UWi9/F/uPZYW6tnWAtfRR6wKBgGsZ\n"
                            + "8MQaCK4JaI/klAhMghgSQnbXZXKVzUZaA3Rln6cX8u7KtgapOOTMlwaZie8Dy1LV\n"
                            + "TTFstAJcm9o3/h1nyYjZy3C4JTUyNpPwqs6enjf7edxVI4eidwFutZD+xcigqHTa\n"
                            + "aikW2atSrZB3fMIjyF7+5meH+hKOqvNiXOty3qn1AoGAZuVxYQy5FVq3YZxzr3Aa\n"
                            + "Am0ShoXTF6QYIbsaUiUGoa/NlHcw9V/lj4AqBRbxbaYMD+hz2J/od9cb268eJKY8\n"
                            + "3b6MvaUqdNhNnWodJXLhgtmGEHDKmTppz2JSTx/tVzCfhFdcOC79StZvcKLhtoFQ\n"
                            + "+/3lEw6NCIXzm5E4+dtJG4k=\n"
                            + "-----END PRIVATE KEY-----\n")
                    .withClientEmail("asc-135@asc-project-1234.iam.gserviceaccount.com")
                    .withClientId("105889053725632919854")
                    .withAuthUri("https://accounts.google.com/o/oauth2/auth")
                    .withTokenUri("https://oauth2.googleapis.com/token")
                    .withAuthProviderX509CertUrl("https://www.googleapis.com/oauth2/v1/certs")
                    .withClientX509CertUrl(
                        "https://www.googleapis.com/robot/v1/metadata/x509/asc-135%40asc-project-1234.iam.gserviceaccount.com"))
            .create();
    }
}
