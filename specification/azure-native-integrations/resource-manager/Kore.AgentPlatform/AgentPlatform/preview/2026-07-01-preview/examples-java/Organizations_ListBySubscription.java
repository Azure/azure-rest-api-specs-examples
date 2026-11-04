
/**
 * Samples for Organizations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Organizations_ListBySubscription.json
     */
    /**
     * Sample code: Organizations_ListBySubscription.
     * 
     * @param manager Entry point to KoreAIArtemisManager.
     */
    public static void
        organizationsListBySubscription(com.azure.resourcemanager.koreaiartemis.KoreAIArtemisManager manager) {
        manager.organizations().list(com.azure.core.util.Context.NONE);
    }
}
