
/**
 * Samples for Environment Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Environment_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Environment_Delete_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void environmentDeleteMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.environments().delete("rgconfluent", "sowkvcymfiziohnofcudjyyaro", "lnmkjsylkxqqyrqmdaf",
            com.azure.core.util.Context.NONE);
    }
}
