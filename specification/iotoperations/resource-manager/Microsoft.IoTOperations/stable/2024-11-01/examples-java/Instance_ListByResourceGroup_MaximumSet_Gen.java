
/**
 * Samples for Instance ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/Instance_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: Instance_ListByResourceGroup.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        instanceListByResourceGroup(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.instances().listByResourceGroup("rgiotoperations", com.azure.core.util.Context.NONE);
    }
}
