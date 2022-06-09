```java
import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2022-02-14/examples/OperationsList.json
     */
    /**
     * Sample code: Retrieve operations list.
     *
     * @param manager Entry point to ImageBuilderManager.
     */
    public static void retrieveOperationsList(com.azure.resourcemanager.imagebuilder.ImageBuilderManager manager) {
        manager.operations().list(Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-imagebuilder_1.0.0-beta.3/sdk/imagebuilder/azure-resourcemanager-imagebuilder/README.md) on how to add the SDK to your project and authenticate.
