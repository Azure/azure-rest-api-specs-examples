
/**
 * Samples for Images ListByDisconnectedOperation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/Images_ListByDisconnectedOperation_MaximumSet_Gen.json
     */
    /**
     * Sample code: Images_ListByDisconnectedOperation.
     * 
     * @param manager Entry point to DisconnectedOperationsManager.
     */
    public static void imagesListByDisconnectedOperation(
        com.azure.resourcemanager.disconnectedoperations.DisconnectedOperationsManager manager) {
        manager.images().listByDisconnectedOperation("rgdisconnectedoperations", "w_-EG-3-euL7K3-E",
            "toynendoobwkrcwmfdfup", 20, 3, com.azure.core.util.Context.NONE);
    }
}
