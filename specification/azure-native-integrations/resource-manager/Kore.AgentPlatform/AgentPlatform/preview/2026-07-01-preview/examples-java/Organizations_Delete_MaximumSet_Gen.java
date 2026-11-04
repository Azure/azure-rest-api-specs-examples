
/**
 * Samples for Organizations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Organizations_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to KoreAIArtemisManager.
     */
    public static void
        organizationsDeleteMaximumSetGen(com.azure.resourcemanager.koreaiartemis.KoreAIArtemisManager manager) {
        manager.organizations().delete("contoso-rg", "contoso-org", com.azure.core.util.Context.NONE);
    }
}
