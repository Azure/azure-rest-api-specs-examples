
import com.azure.resourcemanager.storage.models.ConnectorUpdate;
import com.azure.resourcemanager.storage.models.DataShareSourceUpdate;
import com.azure.resourcemanager.storage.models.ManagedIdentityAuthPropertiesUpdate;
import com.azure.resourcemanager.storage.models.StorageConnectorPropertiesUpdate;

/**
 * Samples for Connectors Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/StorageConnectorCRUD/StorageConnectors_Update.json
     */
    /**
     * Sample code: UpdateConnector.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void updateConnector(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getConnectors().update("testrg", "teststorageaccount", "testconnector",
            new ConnectorUpdate()
                .withProperties(new StorageConnectorPropertiesUpdate().withSource(new DataShareSourceUpdate()
                    .withAuthProperties(new ManagedIdentityAuthPropertiesUpdate().withIdentityResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/newTestIdentity")))),
            com.azure.core.util.Context.NONE);
    }
}
