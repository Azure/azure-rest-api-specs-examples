
/**
 * Samples for Organizations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Organizations_Delete.json
     */
    /**
     * Sample code: Organizations_Delete.
     * 
     * @param manager Entry point to KoreAIArtemisManager.
     */
    public static void organizationsDelete(com.azure.resourcemanager.koreaiartemis.KoreAIArtemisManager manager) {
        manager.organizations().delete("contoso-rg", "contoso-org", com.azure.core.util.Context.NONE);
    }
}
