/** Samples for ReplicationEligibilityResultsOperation List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationEligibilityResults_List.json
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
            .listWithResponse("testRg1", "testVm2", com.azure.core.util.Context.NONE);
    }
}
