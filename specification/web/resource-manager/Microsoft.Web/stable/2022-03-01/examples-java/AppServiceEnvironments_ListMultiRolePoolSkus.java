import com.azure.core.util.Context;

/** Samples for AppServiceEnvironments ListMultiRolePoolSkus. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/AppServiceEnvironments_ListMultiRolePoolSkus.json
     */
    /**
     * Sample code: Get available SKUs for scaling a multi-role pool.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAvailableSKUsForScalingAMultiRolePool(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getAppServiceEnvironments()
            .listMultiRolePoolSkus("test-rg", "test-ase", Context.NONE);
    }
}
