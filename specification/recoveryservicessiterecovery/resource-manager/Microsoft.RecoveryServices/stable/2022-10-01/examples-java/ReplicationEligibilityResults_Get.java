/** Samples for ReplicationEligibilityResultsOperation Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ReplicationEligibilityResults_Get.json
     */
    /**
     * Sample code: Gets the validation errors in case the VM is unsuitable for protection.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheValidationErrorsInCaseTheVMIsUnsuitableForProtection(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager
            .replicationEligibilityResultsOperations()
            .getWithResponse("testRg1", "testVm1", com.azure.core.util.Context.NONE);
    }
}
