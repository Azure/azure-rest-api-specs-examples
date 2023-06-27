import com.azure.core.util.Context;

/** Samples for CustomDomains List. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/CustomDomains_List.json
     */
    /**
     * Sample code: CustomDomains_List.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void customDomainsList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getCustomDomains()
            .list("myResourceGroup", "myservice", "myapp", Context.NONE);
    }
}
