
import com.azure.resourcemanager.appcontainers.models.ManagedCertificateDomainControlValidation;
import com.azure.resourcemanager.appcontainers.models.ManagedCertificateProperties;

/**
 * Samples for ManagedCertificates CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/ManagedCertificate_CreateOrUpdate.
     * json
     */
    /**
     * Sample code: Create or Update Certificate.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        createOrUpdateCertificate(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.managedCertificates().define("certificate-firendly-name").withRegion("East US")
            .withExistingManagedEnvironment("examplerg", "testcontainerenv")
            .withProperties(new ManagedCertificateProperties().withSubjectName("my-subject-name.company.country.net")
                .withDomainControlValidation(ManagedCertificateDomainControlValidation.CNAME))
            .create();
    }
}
