import com.azure.core.util.Context;

/** Samples for Services ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Services_List.json
     */
    /**
     * Sample code: Services_List.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void servicesList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getServices()
            .listByResourceGroup("myResourceGroup", Context.NONE);
    }
}
