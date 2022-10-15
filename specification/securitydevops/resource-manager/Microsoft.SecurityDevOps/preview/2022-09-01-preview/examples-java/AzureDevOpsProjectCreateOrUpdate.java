import com.azure.resourcemanager.securitydevops.models.AutoDiscovery;
import com.azure.resourcemanager.securitydevops.models.AzureDevOpsProjectProperties;

/** Samples for AzureDevOpsProject CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsProjectCreateOrUpdate.json
     */
    /**
     * Sample code: AzureDevOpsProject_CreateOrUpdate.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void azureDevOpsProjectCreateOrUpdate(
        com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager
            .azureDevOpsProjects()
            .define("myProject")
            .withExistingOrg("westusrg", "testconnector", "myOrg")
            .withProperties(new AzureDevOpsProjectProperties().withAutoDiscovery(AutoDiscovery.DISABLED))
            .create();
    }
}
