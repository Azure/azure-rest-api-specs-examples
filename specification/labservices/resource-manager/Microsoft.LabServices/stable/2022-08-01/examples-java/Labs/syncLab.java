import com.azure.core.util.Context;

/** Samples for Labs SyncGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Labs/syncLab.json
     */
    /**
     * Sample code: syncLab.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void syncLab(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.labs().syncGroup("testrg123", "testlab", Context.NONE);
    }
}
