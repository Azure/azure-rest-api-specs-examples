
import com.azure.resourcemanager.cosmos.models.MongoUserDefinitionCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.Role;
import java.util.Arrays;

/**
 * Samples for MongoDBResources CreateUpdateMongoUserDefinition.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBMongoDBUserDefinitionCreateUpdate.json
     */
    /**
     * Sample code: CosmosDBMongoDBUserDefinitionCreateUpdate.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBMongoDBUserDefinitionCreateUpdate(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getMongoDBResources().createUpdateMongoUserDefinition("myMongoUserDefinitionId",
            "myResourceGroupName", "myAccountName",
            new MongoUserDefinitionCreateUpdateParameters().withUsername("myUserName")
                .withPassword("fakeTokenPlaceholder").withDatabaseName("sales").withCustomData("My custom data")
                .withRoles(Arrays.asList(new Role().withDb("sales").withRole("myReadRole")))
                .withMechanisms("SCRAM-SHA-256"),
            com.azure.core.util.Context.NONE);
    }
}
