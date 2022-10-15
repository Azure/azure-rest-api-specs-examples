import com.azure.core.util.Context;
import com.azure.resourcemanager.securitydevops.models.AutoDiscovery;
import com.azure.resourcemanager.securitydevops.models.AzureDevOpsProject;
import com.azure.resourcemanager.securitydevops.models.AzureDevOpsProjectProperties;

/** Samples for AzureDevOpsProject Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsProjectUpdate.json
     */
    /**
     * Sample code: AzureDevOpsProject_Update.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void azureDevOpsProjectUpdate(
        com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        AzureDevOpsProject resource =
            manager
                .azureDevOpsProjects()
                .getWithResponse("westusrg", "testconnector", "myOrg", "myProject", Context.NONE)
                .getValue();
        resource
            .update()
            .withProperties(new AzureDevOpsProjectProperties().withAutoDiscovery(AutoDiscovery.DISABLED))
            .apply();
    }
}
