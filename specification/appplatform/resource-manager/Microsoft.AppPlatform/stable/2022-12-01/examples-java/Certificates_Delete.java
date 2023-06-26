import com.azure.core.util.Context;

/** Samples for Certificates Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Certificates_Delete.json
     */
    /**
     * Sample code: Certificates_Delete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void certificatesDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getCertificates()
            .delete("myResourceGroup", "myservice", "mycertificate", Context.NONE);
    }
}
