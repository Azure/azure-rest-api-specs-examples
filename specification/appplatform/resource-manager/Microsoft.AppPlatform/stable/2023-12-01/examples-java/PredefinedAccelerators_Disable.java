
/**
 * Samples for PredefinedAccelerators Disable.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * PredefinedAccelerators_Disable.json
     */
    /**
     * Sample code: PredefinedAccelerators_Disable.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void predefinedAcceleratorsDisable(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getPredefinedAccelerators().disable("myResourceGroup",
            "myservice", "default", "acc-name", com.azure.core.util.Context.NONE);
    }
}
