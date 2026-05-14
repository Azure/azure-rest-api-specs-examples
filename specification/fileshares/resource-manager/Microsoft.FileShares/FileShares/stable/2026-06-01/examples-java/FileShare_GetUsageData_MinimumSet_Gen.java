
/**
 * Samples for InformationalOperations GetUsageData.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/FileShare_GetUsageData_MinimumSet_Gen.json
     */
    /**
     * Sample code: FileShare_GetUsageData_MinimumSet.
     * 
     * @param manager Entry point to FileSharesManager.
     */
    public static void fileShareGetUsageDataMinimumSet(com.azure.resourcemanager.fileshares.FileSharesManager manager) {
        manager.informationalOperations().getUsageDataWithResponse("westus", com.azure.core.util.Context.NONE);
    }
}
