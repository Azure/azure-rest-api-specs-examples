
/**
 * Samples for Organization GetClusterById.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Organization_GetClusterById_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organization_GetClusterById_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void
        organizationGetClusterByIdMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().getClusterByIdWithResponse("rgconfluent", "qiasyqphlvkxxgyofmf", "xmkhyxmtjzez",
            "lirhyplbzq", com.azure.core.util.Context.NONE);
    }
}
