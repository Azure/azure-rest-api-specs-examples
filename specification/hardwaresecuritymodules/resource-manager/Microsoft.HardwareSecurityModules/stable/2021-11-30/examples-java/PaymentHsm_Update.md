Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-hardwaresecuritymodules_1.0.0-beta.1/sdk/hardwaresecuritymodules/azure-resourcemanager-hardwaresecuritymodules/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.hardwaresecuritymodules.models.DedicatedHsm;
import java.util.HashMap;
import java.util.Map;

/** Samples for DedicatedHsm Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/stable/2021-11-30/examples/PaymentHsm_Update.json
     */
    /**
     * Sample code: Update an existing payment HSM.
     *
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void updateAnExistingPaymentHSM(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        DedicatedHsm resource =
            manager.dedicatedHsms().getByResourceGroupWithResponse("hsm-group", "hsm1", Context.NONE).getValue();
        resource.update().withTags(mapOf("Dept", "hsm", "Environment", "dogfood", "Slice", "A")).apply();
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
```
