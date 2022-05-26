Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-confidentialledger_1.0.0-beta.1/sdk/confidentialledger/azure-resourcemanager-confidentialledger/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Ledger ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/stable/2022-05-13/examples/ConfidentialLedger_List.json
     */
    /**
     * Sample code: ConfidentialLedgerList.
     *
     * @param manager Entry point to ConfidentialLedgerManager.
     */
    public static void confidentialLedgerList(
        com.azure.resourcemanager.confidentialledger.ConfidentialLedgerManager manager) {
        manager.ledgers().listByResourceGroup("DummyResourceGroupName", null, Context.NONE);
    }
}
```
