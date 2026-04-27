
/**
 * Samples for UpdateRuns List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/ListUpdateRuns.json
     */
    /**
     * Sample code: List Update runs under cluster resource.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        listUpdateRunsUnderClusterResource(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.updateRuns().list("testrg", "testcluster", "Microsoft4.2203.2.32", com.azure.core.util.Context.NONE);
    }
}
