Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-appcontainers_1.0.0-beta.2/sdk/appcontainers/azure-resourcemanager-appcontainers/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.appcontainers.models.CertificateProperties;

/** Samples for Certificates CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/Certificate_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or Update Certificate.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void createOrUpdateCertificate(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .certificates()
            .define("certificate-firendly-name")
            .withRegion("East US")
            .withExistingManagedEnvironment("examplerg", "testcontainerenv")
            .withProperties(
                new CertificateProperties()
                    .withPassword("private key password")
                    .withValue("PFX-or-PEM-blob".getBytes()))
            .create();
    }
}
```
