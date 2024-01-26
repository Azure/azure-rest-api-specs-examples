
/**
 * Samples for PredefinedAccelerators Enable.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * PredefinedAccelerators_Enable.json
     */
    /**
     * Sample code: PredefinedAccelerators_Enable.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void predefinedAcceleratorsEnable(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getPredefinedAccelerators().enable("myResourceGroup",
            "myservice", "default", "acc-name", com.azure.core.util.Context.NONE);
    }
}
