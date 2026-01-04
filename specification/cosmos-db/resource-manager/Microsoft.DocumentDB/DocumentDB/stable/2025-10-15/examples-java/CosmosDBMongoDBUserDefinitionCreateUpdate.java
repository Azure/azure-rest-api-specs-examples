
import com.azure.resourcemanager.cosmos.models.MongoUserDefinitionCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.Role;
import java.util.Arrays;

/**
 * Samples for MongoDBResources CreateUpdateMongoUserDefinition.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBMongoDBUserDefinitionCreateUpdate.json
     */
    /**
     * Sample code: CosmosDBMongoDBUserDefinitionCreateUpdate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBMongoDBUserDefinitionCreateUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getMongoDBResources().createUpdateMongoUserDefinition(
            "myMongoUserDefinitionId", "myResourceGroupName", "myAccountName",
            new MongoUserDefinitionCreateUpdateParameters().withUsername("myUserName")
                .withPassword("fakeTokenPlaceholder").withDatabaseName("sales").withCustomData("My custom data")
                .withRoles(Arrays.asList(new Role().withDb("sales").withRole("myReadRole")))
                .withMechanisms("SCRAM-SHA-256"),
            com.azure.core.util.Context.NONE);
    }
}
