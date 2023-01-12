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
                com.azure.core.util.Context.NONE);
    }
}
