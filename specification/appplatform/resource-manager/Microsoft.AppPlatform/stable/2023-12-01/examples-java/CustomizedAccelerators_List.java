
/**
 * Samples for CustomizedAccelerators List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * CustomizedAccelerators_List.json
     */
    /**
     * Sample code: CustomizedAccelerators_List.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void customizedAcceleratorsList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getCustomizedAccelerators().list("myResourceGroup",
            "myservice", "default", com.azure.core.util.Context.NONE);
    }
}
