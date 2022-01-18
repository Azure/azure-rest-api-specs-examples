Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-labservices_1.0.0-beta.2/sdk/labservices/azure-resourcemanager-labservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for VirtualMachines Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/VirtualMachines/startVirtualMachine.json
     */
    /**
     * Sample code: startVirtualMachine.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void startVirtualMachine(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.virtualMachines().start("testrg123", "testlab", "template", Context.NONE);
    }
}
```
