
/**
 * Samples for InformationalOperations GetLimits.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/FileShare_GetLimits_MaximumSet_Gen.json
     */
    /**
     * Sample code: FileShare_GetLimits_MaximumSet.
     * 
     * @param manager Entry point to FileSharesManager.
     */
    public static void fileShareGetLimitsMaximumSet(com.azure.resourcemanager.fileshares.FileSharesManager manager) {
        manager.informationalOperations().getLimitsWithResponse("westus", com.azure.core.util.Context.NONE);
    }
}
