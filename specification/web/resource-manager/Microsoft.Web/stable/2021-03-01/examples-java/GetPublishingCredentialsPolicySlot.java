import com.azure.core.util.Context;

/** Samples for WebApps GetFtpAllowedSlot. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetPublishingCredentialsPolicySlot.json
     */
    /**
     * Sample code: Get FTP Allowed.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getFTPAllowed(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getWebApps()
            .getFtpAllowedSlotWithResponse("rg", "testSite", "stage", Context.NONE);
    }
}
