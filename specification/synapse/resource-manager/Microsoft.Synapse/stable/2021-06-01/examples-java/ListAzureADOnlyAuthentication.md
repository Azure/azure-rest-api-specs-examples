Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.5/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for AzureADOnlyAuthentications List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListAzureADOnlyAuthentication.json
     */
    /**
     * Sample code: Get a list of Azure Active Directory Only Authentication property.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getAListOfAzureActiveDirectoryOnlyAuthenticationProperty(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.azureADOnlyAuthentications().list("workspace-6852", "workspace-2080", Context.NONE);
    }
}
```
