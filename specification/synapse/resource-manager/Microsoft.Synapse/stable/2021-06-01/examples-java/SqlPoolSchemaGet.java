import com.azure.core.util.Context;

/** Samples for SqlPoolSchemas Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SqlPoolSchemaGet.json
     */
    /**
     * Sample code: Get database schema.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getDatabaseSchema(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.sqlPoolSchemas().getWithResponse("myRG", "serverName", "myDatabase", "dbo", Context.NONE);
    }
}
