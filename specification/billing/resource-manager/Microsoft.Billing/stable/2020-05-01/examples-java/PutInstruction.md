Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-billing_1.0.0-beta.2/sdk/billing/azure-resourcemanager-billing/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.billing.fluent.models.InstructionInner;
import java.time.OffsetDateTime;

/** Samples for Instructions Put. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/PutInstruction.json
     */
    /**
     * Sample code: PutInstruction.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void putInstruction(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .instructions()
            .putWithResponse(
                "{billingAccountName}",
                "{billingProfileName}",
                "{instructionName}",
                new InstructionInner()
                    .withAmount(5000f)
                    .withStartDate(OffsetDateTime.parse("2019-12-30T21:26:47.997Z"))
                    .withEndDate(OffsetDateTime.parse("2020-12-30T21:26:47.997Z")),
                Context.NONE);
    }
}
```
