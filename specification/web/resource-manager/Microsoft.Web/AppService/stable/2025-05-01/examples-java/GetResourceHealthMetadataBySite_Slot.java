
/**
 * Samples for ResourceHealthMetadata GetBySiteSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetResourceHealthMetadataBySite_Slot.json
     */
    /**
     * Sample code: Get ResourceHealthMetadata.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getResourceHealthMetadata(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getResourceHealthMetadatas().getBySiteSlotWithResponse("Default-Web-NorthCentralUS",
            "newsiteinnewASE-NCUS", "Production", com.azure.core.util.Context.NONE);
    }
}
