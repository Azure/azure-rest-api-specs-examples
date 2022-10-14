import com.azure.core.util.Context;
import com.azure.resourcemanager.securitydevops.models.AzureDevOpsRepo;

/** Samples for AzureDevOpsRepo Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsRepoUpdate.json
     */
    /**
     * Sample code: AzureDevOpsRepo_Update.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void azureDevOpsRepoUpdate(com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        AzureDevOpsRepo resource =
            manager
                .azureDevOpsRepoes()
                .getWithResponse("westusrg", "testconnector", "myOrg", "myProject", "myRepo", Context.NONE)
                .getValue();
        resource.update().apply();
    }
}
