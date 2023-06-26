import com.azure.core.util.Context;

/** Samples for Certificates Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Certificates_Get.json
     */
    /**
     * Sample code: Certificates_Get.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void certificatesGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getCertificates()
            .getWithResponse("myResourceGroup", "myservice", "mycertificate", Context.NONE);
    }
}
