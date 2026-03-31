
/**
 * Samples for ResourceHealthMetadata ListBySiteSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListResourceHealthMetadataBySite_Slot.json
     */
    /**
     * Sample code: List ResourceHealthMetadata for a site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        listResourceHealthMetadataForASite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getResourceHealthMetadatas().listBySiteSlot("Default-Web-NorthCentralUS",
            "newsiteinnewASE-NCUS", "Production", com.azure.core.util.Context.NONE);
    }
}
