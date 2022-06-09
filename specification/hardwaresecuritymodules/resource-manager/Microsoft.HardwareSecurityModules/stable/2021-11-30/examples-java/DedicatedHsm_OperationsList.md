```java
import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/stable/2021-11-30/examples/DedicatedHsm_OperationsList.json
     */
    /**
     * Sample code: Get a list of Dedicated HSM operations.
     *
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void getAListOfDedicatedHSMOperations(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.operations().list(Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-hardwaresecuritymodules_1.0.0-beta.1/sdk/hardwaresecuritymodules/azure-resourcemanager-hardwaresecuritymodules/README.md) on how to add the SDK to your project and authenticate.
