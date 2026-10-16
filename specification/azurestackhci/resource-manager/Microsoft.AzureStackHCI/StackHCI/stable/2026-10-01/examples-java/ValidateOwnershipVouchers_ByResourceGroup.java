
import com.azure.resourcemanager.azurestackhci.models.OwnerKeyType;
import com.azure.resourcemanager.azurestackhci.models.OwnershipVoucherDetails;
import com.azure.resourcemanager.azurestackhci.models.ValidateOwnershipVouchersRequest;
import java.util.Arrays;

/**
 * Samples for OwnershipVouchers Validate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/ValidateOwnershipVouchers_ByResourceGroup.json
     */
    /**
     * Sample code: Validate ownership vouchers in a given resource group.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void validateOwnershipVouchersInAGivenResourceGroup(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.ownershipVouchers().validateWithResponse("ArcInstance-rg", "westus",
            new ValidateOwnershipVouchersRequest().withOwnershipVoucherDetails(
                Arrays.asList(new OwnershipVoucherDetails().withOwnershipVoucher("Device Model Ownership content")
                    .withOwnerKeyType(OwnerKeyType.MICROSOFT_MANAGED))),
            com.azure.core.util.Context.NONE);
    }
}
