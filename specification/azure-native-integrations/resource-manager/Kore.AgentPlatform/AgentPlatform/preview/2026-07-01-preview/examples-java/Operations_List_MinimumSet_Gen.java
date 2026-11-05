
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Operations_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: Operations_List_MinimumSet_Gen.
     * 
     * @param manager Entry point to KoreAIArtemisManager.
     */
    public static void
        operationsListMinimumSetGen(com.azure.resourcemanager.koreaiartemis.KoreAIArtemisManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
