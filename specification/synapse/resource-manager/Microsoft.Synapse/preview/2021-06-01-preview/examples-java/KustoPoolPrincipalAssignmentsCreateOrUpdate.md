Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.4/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.synapse.models.ClusterPrincipalRole;
import com.azure.resourcemanager.synapse.models.PrincipalType;

/** Samples for KustoPoolPrincipalAssignments CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolPrincipalAssignmentsCreateOrUpdate.json
     */
    /**
     * Sample code: KustoPoolPrincipalAssignmentsCreateOrUpdate.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolPrincipalAssignmentsCreateOrUpdate(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolPrincipalAssignments()
            .define("kustoprincipal1")
            .withExistingKustoPool("synapseWorkspaceName", "kustoclusterrptest4", "kustorptest")
            .withPrincipalId("87654321-1234-1234-1234-123456789123")
            .withRole(ClusterPrincipalRole.ALL_DATABASES_ADMIN)
            .withTenantId("12345678-1234-1234-1234-123456789123")
            .withPrincipalType(PrincipalType.APP)
            .create();
    }
}
```
