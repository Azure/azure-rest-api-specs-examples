
import com.azure.resourcemanager.cosmos.models.DatabaseAccountRegenerateKeyParameters;
import com.azure.resourcemanager.cosmos.models.KeyKind;

/**
 * Samples for DatabaseAccounts RegenerateKey.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBDatabaseAccountRegenerateKey.json
     */
    /**
     * Sample code: CosmosDBDatabaseAccountRegenerateKey.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBDatabaseAccountRegenerateKey(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getDatabaseAccounts().regenerateKey("rg1", "ddb1",
            new DatabaseAccountRegenerateKeyParameters().withKeyKind(KeyKind.PRIMARY),
            com.azure.core.util.Context.NONE);
    }
}
