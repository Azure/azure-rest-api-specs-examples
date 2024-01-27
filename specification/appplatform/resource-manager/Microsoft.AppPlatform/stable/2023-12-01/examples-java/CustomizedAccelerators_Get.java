
/**
 * Samples for CustomizedAccelerators Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * CustomizedAccelerators_Get.json
     */
    /**
     * Sample code: CustomizedAccelerators_Get.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void customizedAcceleratorsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getCustomizedAccelerators().getWithResponse("myResourceGroup",
            "myservice", "default", "acc-name", com.azure.core.util.Context.NONE);
    }
}
