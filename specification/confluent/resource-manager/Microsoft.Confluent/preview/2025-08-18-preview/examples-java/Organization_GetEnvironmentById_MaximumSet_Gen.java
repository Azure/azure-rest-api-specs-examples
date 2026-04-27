
/**
 * Samples for Organization GetEnvironmentById.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Organization_GetEnvironmentById_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organization_GetEnvironmentById_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void
        organizationGetEnvironmentByIdMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().getEnvironmentByIdWithResponse("rgconfluent", "p", "kvifvjnmbilj",
            com.azure.core.util.Context.NONE);
    }
}
