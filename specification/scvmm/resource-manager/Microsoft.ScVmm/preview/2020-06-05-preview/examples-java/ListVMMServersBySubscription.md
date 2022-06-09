```java
import com.azure.core.util.Context;

/** Samples for VmmServers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/ListVMMServersBySubscription.json
     */
    /**
     * Sample code: ListVmmServersBySubscription.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void listVmmServersBySubscription(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.vmmServers().list(Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-scvmm_1.0.0-beta.1/sdk/scvmm/azure-resourcemanager-scvmm/README.md) on how to add the SDK to your project and authenticate.
