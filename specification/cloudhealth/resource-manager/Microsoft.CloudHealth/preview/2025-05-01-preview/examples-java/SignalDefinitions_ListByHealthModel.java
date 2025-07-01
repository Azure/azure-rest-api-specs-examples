
/**
 * Samples for SignalDefinitions ListByHealthModel.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/SignalDefinitions_ListByHealthModel.json
     */
    /**
     * Sample code: SignalDefinitions_ListByHealthModel.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void
        signalDefinitionsListByHealthModel(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.signalDefinitions().listByHealthModel("rgopenapi", "myHealthModel", null,
            com.azure.core.util.Context.NONE);
    }
}
