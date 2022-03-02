Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mobilenetwork_1.0.0-beta.1/sdk/mobilenetwork/azure-resourcemanager-mobilenetwork/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for SimPolicies ListByMobileNetwork. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/SimPolicyListByMobileNetwork.json
     */
    /**
     * Sample code: List sim policies in a mobile network.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listSimPoliciesInAMobileNetwork(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.simPolicies().listByMobileNetwork("testResourceGroupName", "testMobileNetwork", Context.NONE);
    }
}
```
