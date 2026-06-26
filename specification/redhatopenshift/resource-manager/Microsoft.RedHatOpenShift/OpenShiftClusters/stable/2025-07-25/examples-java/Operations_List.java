
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-25/Operations_List.json
     */
    /**
     * Sample code: Lists all of the available RP operations.
     * 
     * @param manager Entry point to RedHatOpenShiftManager.
     */
    public static void
        listsAllOfTheAvailableRPOperations(com.azure.resourcemanager.redhatopenshift.RedHatOpenShiftManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
