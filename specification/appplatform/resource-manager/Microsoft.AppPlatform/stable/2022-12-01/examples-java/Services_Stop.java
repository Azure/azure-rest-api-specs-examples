import com.azure.core.util.Context;

/** Samples for Services Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Services_Stop.json
     */
    /**
     * Sample code: Services_Stop.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void servicesStop(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getServices()
            .stop("myResourceGroup", "myservice", Context.NONE);
    }
}
