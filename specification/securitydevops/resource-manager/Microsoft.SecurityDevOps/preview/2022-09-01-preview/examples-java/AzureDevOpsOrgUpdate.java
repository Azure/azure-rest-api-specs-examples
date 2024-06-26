import com.azure.core.util.Context;
import com.azure.resourcemanager.securitydevops.models.AutoDiscovery;
import com.azure.resourcemanager.securitydevops.models.AzureDevOpsOrg;
import com.azure.resourcemanager.securitydevops.models.AzureDevOpsOrgProperties;

/** Samples for AzureDevOpsOrg Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsOrgUpdate.json
     */
    /**
     * Sample code: AzureDevOpsOrg_Update.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void azureDevOpsOrgUpdate(com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        AzureDevOpsOrg resource =
            manager.azureDevOpsOrgs().getWithResponse("westusrg", "testconnector", "myOrg", Context.NONE).getValue();
        resource
            .update()
            .withProperties(new AzureDevOpsOrgProperties().withAutoDiscovery(AutoDiscovery.DISABLED))
            .apply();
    }
}
