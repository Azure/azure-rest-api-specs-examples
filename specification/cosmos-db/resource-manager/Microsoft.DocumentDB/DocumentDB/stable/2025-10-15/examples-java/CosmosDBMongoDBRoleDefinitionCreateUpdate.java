
import com.azure.resourcemanager.cosmos.models.MongoRoleDefinitionCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.Privilege;
import com.azure.resourcemanager.cosmos.models.PrivilegeResource;
import com.azure.resourcemanager.cosmos.models.Role;
import java.util.Arrays;

/**
 * Samples for MongoDBResources CreateUpdateMongoRoleDefinition.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBMongoDBRoleDefinitionCreateUpdate.json
     */
    /**
     * Sample code: CosmosDBMongoDBRoleDefinitionCreateUpdate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBMongoDBRoleDefinitionCreateUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getMongoDBResources().createUpdateMongoRoleDefinition(
            "myMongoRoleDefinitionId", "myResourceGroupName", "myAccountName",
            new MongoRoleDefinitionCreateUpdateParameters().withRoleName("myRoleName").withDatabaseName("sales")
                .withPrivileges(Arrays.asList(
                    new Privilege().withResource(new PrivilegeResource().withDb("sales").withCollection("sales"))
                        .withActions(Arrays.asList("insert", "find"))))
                .withRoles(Arrays.asList(new Role().withDb("sales").withRole("myInheritedRole"))),
            com.azure.core.util.Context.NONE);
    }
}
