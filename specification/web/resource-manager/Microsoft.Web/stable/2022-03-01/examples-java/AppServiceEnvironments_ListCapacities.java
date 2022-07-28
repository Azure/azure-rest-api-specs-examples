import com.azure.core.util.Context;

/** Samples for AppServiceEnvironments ListCapacities. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/AppServiceEnvironments_ListCapacities.json
     */
    /**
     * Sample code: Get the used, available, and total worker capacity an App Service Environment.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getTheUsedAvailableAndTotalWorkerCapacityAnAppServiceEnvironment(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getAppServiceEnvironments()
            .listCapacities("test-rg", "test-ase", Context.NONE);
    }
}
