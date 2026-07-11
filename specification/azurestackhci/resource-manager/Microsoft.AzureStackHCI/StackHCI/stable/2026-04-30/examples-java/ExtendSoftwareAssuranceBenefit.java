
import com.azure.resourcemanager.azurestackhci.models.SoftwareAssuranceChangeRequest;
import com.azure.resourcemanager.azurestackhci.models.SoftwareAssuranceChangeRequestProperties;
import com.azure.resourcemanager.azurestackhci.models.SoftwareAssuranceIntent;

/**
 * Samples for Clusters ExtendSoftwareAssuranceBenefit.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-30/ExtendSoftwareAssuranceBenefit.json
     */
    /**
     * Sample code: Create cluster Identity.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void createClusterIdentity(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.clusters().extendSoftwareAssuranceBenefit("test-rg", "myCluster",
            new SoftwareAssuranceChangeRequest().withProperties(new SoftwareAssuranceChangeRequestProperties()
                .withSoftwareAssuranceIntent(SoftwareAssuranceIntent.ENABLE)),
            com.azure.core.util.Context.NONE);
    }
}
