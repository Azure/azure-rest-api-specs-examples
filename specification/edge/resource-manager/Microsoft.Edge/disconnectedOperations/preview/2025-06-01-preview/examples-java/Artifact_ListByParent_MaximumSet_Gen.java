
/**
 * Samples for Artifacts ListByParent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/Artifact_ListByParent_MaximumSet_Gen.json
     */
    /**
     * Sample code: Artifacts_ListByParent.
     * 
     * @param manager Entry point to DisconnectedOperationsManager.
     */
    public static void
        artifactsListByParent(com.azure.resourcemanager.disconnectedoperations.DisconnectedOperationsManager manager) {
        manager.artifacts().listByParent("rgdisconnectedoperations", "XOn_Y-7_M-46E-Y", "2v5Q3mNihPV88C882LnbQO8",
            com.azure.core.util.Context.NONE);
    }
}
