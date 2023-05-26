/** Samples for ResourceHealthMetadata GetBySiteSlot. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/GetResourceHealthMetadataBySite.json
     */
    /**
     * Sample code: Get ResourceHealthMetadata.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getResourceHealthMetadata(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getResourceHealthMetadatas()
            .getBySiteSlotWithResponse(
                "Default-Web-NorthCentralUS", "newsiteinnewASE-NCUS", "Production", com.azure.core.util.Context.NONE);
    }
}
