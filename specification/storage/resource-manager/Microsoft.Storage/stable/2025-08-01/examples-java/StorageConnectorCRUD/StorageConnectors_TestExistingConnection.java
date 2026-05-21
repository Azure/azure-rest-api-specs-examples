
import com.azure.resourcemanager.storage.models.TestExistingConnectionRequest;

/**
 * Samples for Connectors TestExistingConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/StorageConnectorCRUD/StorageConnectors_TestExistingConnection.json
     */
    /**
     * Sample code: ExistingConnectionTest.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void existingConnectionTest(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getConnectors().testExistingConnection("testrg", "teststorageaccount", "testconnector",
            new TestExistingConnectionRequest().withUniqueId("12345678-1234-1234-1234-12345678912"),
            com.azure.core.util.Context.NONE);
    }
}
